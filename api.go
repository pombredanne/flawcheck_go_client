type Params struct {
    Count int `url:"count,omitempty"`
}
params := &Params{Count: 5}

req, err := sling.New().Get("https://example.com").QueryStruct(params).Request()
client.Do(req)