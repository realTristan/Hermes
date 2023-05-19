# Hermes ![Stars](https://img.shields.io/github/stars/realTristan/Hermes?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/Hermes?label=Watchers)
![banner](https://user-images.githubusercontent.com/75189508/230987049-665418b1-3576-49b7-861e-29036859ad8a.png)

# CLI
## Usage
```
hermes serve -p 6000
```

## Install
```
MacOS:
  $ curl "https://github.com/realTristan/Hermes/raw/master/cli/hermes" -o /usr/local/bin/hermes
  
Windows:
  $ curl "https://github.com/realTristan/Hermes/raw/master/cli/hermes.exe" -o C:\hermes.exe
  $ set PATH=%PATH%;C:\hermes.exe
```

# Functions
```py
# Wrapper for the Hermes cache
class Hermes:
    # Set a value in the cache
    def set(self, key: str, value: dict, full_text: bool) -> any:
    
    # Get a value from the cache
    def get(self, key: str) -> any:
    
    # Delete a value from the cache
    def delete(self, key: str) -> any:
    
    # Get all keys in the cache
    def keys(self) -> any:
    
    # Get all values in the cache
    def values(self) -> any:
    
    # Get the cache length
    def length(self) -> any:
    
    # Clear the cache
    def clean(self) -> any:
    
    # Get the cache info
    def info(self) -> any:
    
    # Check if value exists in the cache
    def exists(self, key: str) -> any:
    
    # Intialize the full text cache
    def ft_init(self) -> any:
    
    # Clean the full text cache
    def ft_clean(self) -> any:
    
    # Search the full text cache
    def ft_search(self, query: str, strict: bool, limit: int, schema: dict) -> any:
    
    # Search one word in the full text cache
    def ft_search_one(self, query: str, strict: bool, limit: int) -> any:
    
    # Search value in the full text cache
    def ft_search_value(self, query: str, limit: int, schema: dict) -> any:
    
    # Search values with a key in the full text cache
    def ft_search_key(self, query: str, key: str, limit: int) -> any:
    
    # Set the max bytes of the full text cache
    def ft_set_max_bytes(self, max_bytes: int) -> any:
    
    # Set the max words of the full text cache
    def ft_set_max_words(self, max_words: int) -> any:
    
    # Get the full text cache
    def ft_cache(self) -> any:
    
    # Get the full text cache size
    def ft_size(self) -> any:
    
    # Get whether the full text cache is initialized
    def ft_initialized(self) -> any:
    
    # Add to the full text cache
    def ft_add(self, key: str) -> any:
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