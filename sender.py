import grpc
import time
import sys
import os

current_dir = os.path.dirname(os.path.abspath(__file__))
grpc_gen_dir = os.path.join(current_dir, 'grpc', 'gen')
sys.path.append(grpc_gen_dir)

import service_pb2
import service_pb2_grpc

def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = service_pb2_grpc.StringServiceStub(channel)

    string_message = service_pb2.StringMessage(content="Hey, Explain GoLang to me in 2 sentences.")
    stub.SendString(string_message)
    print("Sent message:", string_message.content)

if __name__ == '__main__':
    run()
