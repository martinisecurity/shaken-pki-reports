# STIR/SHAKEN Certificate Repository Compliance

## Rogers Communications Inc

Name: `https://sticr-cstga.ccid.neustar/api/v1/certificate/a743a54c63164030723c14d15579fcca.pem`\
Tested At: 26 Aug 24 17:59 UTC\
Time: 318ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:03 UTC