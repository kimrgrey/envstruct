# envstruct

Library that allows to work with environment variables using structs in Go. 

## Usage

Using reflection `envstruct` is able to automatically load values of environment variables based on field name in your 
structure, so you can just move from this:

```go
type Config struct {
  Port string
  Host string
}

func main() {
  c := Config{}
  c.Port = os.Getenv("PORT")
  c.Host = os.Getenv("HOST")
  // ...
}
```

to this:

```go
type Config struct {
  Port string
  Host string
}

func main() {
  config := &Config{}
  envstruct.Load(config)
  // ...
}
```

## License

MIT License

Copyright (c) 2017 Sergey Tsvetkov

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
