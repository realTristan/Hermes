# Hermes
Extremely Fast Full-Text-Search Algorithm

# About
## Storing Data
Hermes works by iterating over the items in the data.json file, and then iterates over the keys and values of the items and splits the value into different words. It then stores the indices for all of the items that contain those words in a dictionary.

## Accessing Data
When searching for a word, Hermes will return a list of indices for all of the items that contain that word. It checks whether the key in the cache dictionary contains the provided word, instead of just accessing it so that short forms for words can be used.

# Example
```py
from fastapi.middleware.cors import CORSMiddleware
from fastapi import FastAPI, Request
from cache import Cache
import time

# // The FastAPI app
app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# // The cache
cache: Cache = Cache()
cache.load()

# // Courses endpoint
@app.get("/courses")
async def root(request: Request):
    # // Get the course to search for from the query params
    course: str = request.query_params.get("q", "CS")

    # // Search for a word in the cache
    start_time: float = time.time()
    indices: list[int] = cache.search(course)
    print(f"Search time: {time.time() - start_time} seconds")

    # // Convert the indices to the actual items
    items: list[dict] = cache.indices_to_data(indices)
    print(f"Conversion time: {time.time() - start_time} seconds")

    # // Return the items
    return items

# // Run the app
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="localhost", port=8000)
```

# License
MIT License

Copyright (c) 2023 Tristan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.