//Execute a query that will automatically be throttled
func (c *TwitterApi) throttledQuery(queryQueue chan queryChan) {
    for q := range queryQueue {

        endpoint_path := q.endpoint_path
        method := q.method
        keyVals := q.keyVals
        response_ch := q.response_ch
        err := c.execQuery(url, form, data, method)
        response_ch <- struct {
            result []byte
            err    error
        }{result, err}

        // Reading from this will block until the specified time has elapsed
        <-time.After(SECONDS_PER_QUERY)
    }
}
