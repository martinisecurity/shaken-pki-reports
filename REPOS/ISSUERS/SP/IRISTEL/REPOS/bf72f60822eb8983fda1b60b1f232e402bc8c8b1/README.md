# STIR/SHAKEN Certificate Repository Compliance

## IRISTEL

Name: `https://sticr-cstga.ccid.neustar/api/v1/certificate/d65e891e9a73bd32043540a43bc0908a.pem`\
Tested At: 18 Aug 25 21:08 UTC\
Time: 301ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC