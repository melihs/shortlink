import { check } from 'k6'
import http from 'k6/http'
import tracing from "k6/experimental/tracing"

// eslint-disable-next-line import/no-mutable-exports
export let options = {
  ext: {
    loadimpact: {
      // eslint-disable-next-line no-undef
      projectID: __ENV.K6_PROJECT_ID,
      // Test runs with the same name groups test runs together
      name: 'HTTP main page',
    },
  },
}

tracing.instrumentHTTP({
  // possible values: "w3c", "jaeger"
  propagator: "w3c",
})

export default () => {
  // eslint-disable-next-line no-undef
  const TARGET_HOSTNAME = __ENV.TARGET_HOSTNAME || 'localhost:3001'

  const res = http.get(`http://${TARGET_HOSTNAME}/`, {
    headers: {
      'trace-id': 'instrumented/get',
    },
  })
  check(res, {
    'is status 200': (r) => r.status === 200,
  })
}
