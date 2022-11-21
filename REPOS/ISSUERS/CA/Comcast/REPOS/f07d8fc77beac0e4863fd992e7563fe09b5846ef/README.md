# STIR/SHAKEN Certificate Repository Compliance

## Comcast

Name: `https://sticr.stir.comcast.com/46330114-1c23-4d81-85c1-c74f8a443aec.cer`\
Tested At: 21 Nov 22 20:33 UTC\
Time: 19ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Nov 22 20:33 UTC