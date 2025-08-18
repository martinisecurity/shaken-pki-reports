# STIR/SHAKEN Certificate Repository Compliance

## Microtalk Europe Limited

Name: `https://appreg.telcoportal.com/mobileapps/neustar23/9c680c2b2d89e44b0a235be6685b8d.cer`\
Tested At: 18 Aug 25 20:04 UTC\
Time: 336ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC