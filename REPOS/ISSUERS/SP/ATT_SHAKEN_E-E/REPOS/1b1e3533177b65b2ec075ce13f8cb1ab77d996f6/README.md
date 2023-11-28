# STIR/SHAKEN Certificate Repository Compliance

## ATT SHAKEN E-E

Name: `https://cert2.sticr.att.net:8443/sti-cr/att-stica1661870921294-cert.crt`\
Tested At: 28 Nov 23 15:59 UTC\
Time: 153ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 28 Nov 23 16:15 UTC