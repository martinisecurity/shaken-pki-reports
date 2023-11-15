# STIR/SHAKEN Certificate Repository Compliance

## Comcast

Name: `https://sticr.stir.comcast.com/d926b087-09a5-4b97-b21c-90e5e3a0863c.cer`\
Tested At: 15 Nov 23 18:09 UTC\
Time: 115ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 15 Nov 23 18:10 UTC