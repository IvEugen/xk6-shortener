import http from 'k6/http';
import { check } from 'k6';
import { sleep } from 'k6';
import short from 'k6/x/shortener';

export const options = {
  vus: 1,
  duration: '15s',
};

export default function () {
  let jsonDate = short.GeneratorURL(10)
  console.log(`jsonData: ${jsonDate}`);

/*     const res = http.get('http://localhost:8081/api/orders/6642691858');
  check(res, {
    'is status 200': (r) => r.status === 200,
  }); */
  sleep(1);
}
