# STIR/SHAKEN Certificate Repository Compliance

## Unknown

Name: `https://187.174.67.118:8080/7075515eb2d150fc98c43e794c07bbca.cer`\
Tested At: 28 Nov 23 19:52 UTC\
Time: 190ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_http_status_200](../../ISSUES/e_http_status_200/README.md) | error | HTTP | HTTP response shall have StatusCode 200, but it is 404 Not Found |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://187.174.67.118:8080/7075515eb2d150fc98c43e794c07bbca.cer": x509: cannot validate certificate for 187.174.67.118 because it doesn't contain any IP SANs |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |
| [w_atis_protocol](../../ISSUES/w_atis_protocol/README.md) | warn | ATIS-1000080 | The verifier should not dereference any protocol other than https or a port other than 443 or 8443 |

Generated: 28 Nov 23 20:21 UTC