# STIR/SHAKEN Certificate Repository Compliance

## NetNumber

| Code | Source | Instances |
|------|--------|-----------|
| [e_atis_cache_header](../ISSUES/e_atis_cache_header/README.md) | ATIS-1000074 | 2 |
| [w_atis_content_type](../ISSUES/w_atis_content_type/README.md) | ATIS-1000080 | 2 |

### https://bwt.io/8697a2.crt

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |

### https://certificate.zt.plivo.com/cert09062023.crt

| Code | Status | Source | Details |
|------|--------|--------|---------|
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

### https://www.gstatic.com/gtp/stir/Pkm8HyNNMEiDhJ67_a-tWw.pem

| Code | Status | Source | Details |
|------|--------|--------|---------|

5 tests were ran and no warning or error level issues were found

### https://www.gstatic.com/gtp/stir/sFCvi_7snpTdPyz7VUUx8g.pem

| Code | Status | Source | Details |
|------|--------|--------|---------|

5 tests were ran and no warning or error level issues were found

### https://www.gstatic.com/gtp/stir/yWm5JCvzCtTKhZihGKFIFA.pem

| Code | Status | Source | Details |
|------|--------|--------|---------|

5 tests were ran and no warning or error level issues were found


Generated: 27/10/2022 at 22:42:50