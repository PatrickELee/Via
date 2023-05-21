from concurrent import futures

import grpc
import via_pb2
import via_pb2_grpc
import model_wrapper as ml

import googlemaps
import os
from dotenv import load_dotenv

class ViaServer(via_pb2_grpc.ViaServicer):
  def __init__(self):
    load_dotenv()
    gmaps_api_key = os.getenv('GMP_API_KEY')
    self.client = googlemaps.Client(key=gmaps_api_key)  
  
  def GetDangerProbability(self, request, context):
    res = self.client.geocode(request.location)
    lat = float(res[0]['geometry']['location']['lat'])
    long = float(res[0]['geometry']['location']['lng'])
    time = int(request.time)

    danger_prob = ml.run_knn_model(latitude=lat,longitude=long,time=float(request.time))
    
    danger = 'Low'
    if danger_prob[0] > 0.8 or danger_prob[1] > 0.8:
      danger = 'High'
    
    return via_pb2.GetDangerProbabilityResponse(danger=danger)
	  
def main():
   server = grpc.server(futures.ThreadPoolExecutor(max_workers=2))
   via_pb2_grpc.add_ViaServicer_to_server(ViaServer(), server)
   server.add_insecure_port('[::]:50051')
   print("gRPC starting")
   server.start()
   server.wait_for_termination()

if __name__ == '__main__':
  main()
