# STIR/SHAKEN Certificate Repository Compliance

## GoTo Communications Inc

Name: `https://pstn-cdn.live.gtc.goto.com/certs/stirshaken/goto-2022-09`\
Tested At: 17 Nov 22 19:11 UTC\
Time: 210ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 17 Nov 22 19:21 UTC