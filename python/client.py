import grpc

import via_pb2
import via_pb2_grpc

def run():
   metadata = [('content-type', 'application/grpc')]
   with grpc.insecure_channel('localhost:32314') as channel:
      stub = via_pb2_grpc.ViaStub(channel)
      response = stub.GetDangerProbability(via_pb2.GetDangerProbabilityRequest(lat='43.2341234', long = "26.345433", time="1920"), metadata=metadata)
   print("Greeter client received following from server: " + response.danger) 
  
run()
