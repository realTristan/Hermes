import websockets, asyncio, base64, json, time

def encode(value):
    return base64.b64encode(value.encode("utf-8")).decode("utf-8")

def decode(value):
    return base64.b64decode(value.encode("utf-8")).decode("utf-8")

def test_set():
    value = {
        "test": "test"
    }
    return json.dumps({
        "function": "set",
        "key": "test",
        "value": encode(json.dumps(value)),
        "ft": False
    })

def test_get():
    return json.dumps({
        "function": "get",
        "key": "test"
    })

# connect to wss://127.0.0.1:3000/ws/hermes/cache
async def test():
    async with websockets.connect("ws://127.0.0.1:3000/ws/hermes/cache") as websocket:
        # track start time
        start = time.time()

        # test set
        await websocket.send(test_set())
        print(await websocket.recv())

        # print time taken
        print("Time taken: " + str(time.time() - start))

        # test get
        await websocket.send(test_get())
        print(await websocket.recv())

        # close socket
        await websocket.close()

asyncio.run(test())