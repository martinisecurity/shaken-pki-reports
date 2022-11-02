# STIR/SHAKEN Certificate Repository Compliance

## LiveVox

Name: `https://stir.na3.livevox.com/cert/2B6FU4qN`\
Tested At: 02 Nov 22 15:15 UTC\
Time: 348ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 02 Nov 22 15:15 UTC