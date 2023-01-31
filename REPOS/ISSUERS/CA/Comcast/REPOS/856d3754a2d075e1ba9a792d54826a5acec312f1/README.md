# STIR/SHAKEN Certificate Repository Compliance

## Comcast

Name: `https://sticr.stir.comcast.com/cc1657f6-11b6-4bff-a484-42e34f46c624.cer`\
Tested At: 31 Jan 23 21:50 UTC\
Time: 45ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 31 Jan 23 21:50 UTC