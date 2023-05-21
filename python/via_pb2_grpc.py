# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import via_pb2 as via__pb2


class ViaStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetDangerProbability = channel.unary_unary(
                '/viapb.Via/GetDangerProbability',
                request_serializer=via__pb2.GetDangerProbabilityRequest.SerializeToString,
                response_deserializer=via__pb2.GetDangerProbabilityResponse.FromString,
                )
        self.GetIncident = channel.unary_unary(
                '/viapb.Via/GetIncident',
                request_serializer=via__pb2.GetIncidentRequest.SerializeToString,
                response_deserializer=via__pb2.GetIncidentResponse.FromString,
                )


class ViaServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetDangerProbability(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetIncident(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ViaServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetDangerProbability': grpc.unary_unary_rpc_method_handler(
                    servicer.GetDangerProbability,
                    request_deserializer=via__pb2.GetDangerProbabilityRequest.FromString,
                    response_serializer=via__pb2.GetDangerProbabilityResponse.SerializeToString,
            ),
            'GetIncident': grpc.unary_unary_rpc_method_handler(
                    servicer.GetIncident,
                    request_deserializer=via__pb2.GetIncidentRequest.FromString,
                    response_serializer=via__pb2.GetIncidentResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'viapb.Via', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Via(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetDangerProbability(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/viapb.Via/GetDangerProbability',
            via__pb2.GetDangerProbabilityRequest.SerializeToString,
            via__pb2.GetDangerProbabilityResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetIncident(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/viapb.Via/GetIncident',
            via__pb2.GetIncidentRequest.SerializeToString,
            via__pb2.GetIncidentResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)