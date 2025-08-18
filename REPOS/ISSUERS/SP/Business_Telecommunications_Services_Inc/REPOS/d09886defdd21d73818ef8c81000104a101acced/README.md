# STIR/SHAKEN Certificate Repository Compliance

## Business Telecommunications Services Inc

Name: `https://stirshaken-m.bts.io/api/v1/certificates/5aa4c5cea59d7ff743f94f88b2b149ed.cer`\
Tested At: 18 Aug 25 21:09 UTC\
Time: 254ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC