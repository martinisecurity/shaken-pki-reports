# STIR/SHAKEN Certificate Repository Compliance

## Microtalk USA Inc

Name: `https://appreg.telcoportal.com/mobileapps/neustar23/9c680c2b2d89e44b0a235be6685b8d.cer`\
Tested At: 28 Nov 23 19:52 UTC\
Time: 407ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 28 Nov 23 20:21 UTC