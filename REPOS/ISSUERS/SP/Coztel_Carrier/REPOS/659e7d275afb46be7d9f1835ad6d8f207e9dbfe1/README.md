# STIR/SHAKEN Certificate Repository Compliance

## Coztel Carrier

Name: `https://coztel.com/certificates/191c4c42dd7fa6115e84100637e42c99.cer`\
Tested At: 18 Aug 25 20:16 UTC\
Time: 284ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC