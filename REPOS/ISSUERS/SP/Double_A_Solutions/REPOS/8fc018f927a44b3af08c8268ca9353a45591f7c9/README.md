# STIR/SHAKEN Certificate Repository Compliance

## Double A Solutions

Name: `https://sip.alinephonesystems.com/stirshaken/644K.chain.crt.pem`\
Tested At: 05 Apr 24 18:56 UTC\
Time: 253ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 05 Apr 24 19:04 UTC