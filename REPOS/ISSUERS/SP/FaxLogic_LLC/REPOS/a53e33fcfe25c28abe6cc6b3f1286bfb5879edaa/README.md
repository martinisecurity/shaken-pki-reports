# STIR/SHAKEN Certificate Repository Compliance

## FaxLogic LLC

Name: `https://cdn.pgxn.net/sti/20230701.cer`\
Tested At: 28 Nov 22 20:37 UTC\
Time: 124ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 28 Nov 22 20:41 UTC