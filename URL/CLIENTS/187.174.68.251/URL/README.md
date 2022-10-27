# STIR/SHAKEN Certificate Repository Compliance

## 187.174.68.251

| Code | Source | Instances |
|------|--------|-----------|
| [w_atis_protocol](../ISSUES/w_atis_protocol/README.md) | ATIS-1000080 | 1 |
| [e_atis_cache_header](../ISSUES/e_atis_cache_header/README.md) | ATIS-1000074 | 1 |
| [w_atis_content_type](../ISSUES/w_atis_content_type/README.md) | ATIS-1000080 | 1 |
| [e_tls_transport](../ISSUES/e_tls_transport/README.md) | System | 1 |

### https://187.174.68.251:8080/7075515eb2d150fc98c43e794c07bbca.cer

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_atis_protocol | warn | ATIS-1000080 | The verifier should not dereference any protocol other than https or a port other than 443 or 8443 |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| e_tls_transport | error | System | Get "https://187.174.68.251:8080/7075515eb2d150fc98c43e794c07bbca.cer": x509: cannot validate certificate for 187.174.68.251 because it doesn't contain any IP SANs |


Generated: 27/10/2022 at 22:42:50