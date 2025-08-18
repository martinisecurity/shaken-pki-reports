# STIR/SHAKEN Certificate Repository Compliance

## Ringover Inc

Name: `https://385e4881-f5b7-494a-8dc5-1367ee585926.s3.amazonaws.com/265a1907e1c43af84ed684b4fae9bd52.cer`\
Tested At: 18 Aug 25 20:04 UTC\
Time: 123ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC