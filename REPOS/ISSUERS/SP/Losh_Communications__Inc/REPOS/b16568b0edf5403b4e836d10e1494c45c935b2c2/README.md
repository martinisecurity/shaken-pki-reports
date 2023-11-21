# STIR/SHAKEN Certificate Repository Compliance

## Losh Communications, Inc

Name: `https://nyc01.trunks2.calldecibel.com:5000/stirshaken_certs/149K.pem`\
Tested At: 21 Nov 23 19:09 UTC\
Time: 443ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_protocol](../../ISSUES/w_atis_protocol/README.md) | warn | ATIS-1000080 | The verifier should not dereference any protocol other than https or a port other than 443 or 8443 |

Generated: 21 Nov 23 19:18 UTC