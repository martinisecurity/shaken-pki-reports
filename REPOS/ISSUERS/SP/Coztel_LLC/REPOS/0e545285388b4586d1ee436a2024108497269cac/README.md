# STIR/SHAKEN Certificate Repository Compliance

## Coztel LLC

Name: `https://coztel.com/certificates/319aa14558f159d3c8c81427198aad03.cer`\
Tested At: 26 Aug 24 18:11 UTC\
Time: 279ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:49 UTC