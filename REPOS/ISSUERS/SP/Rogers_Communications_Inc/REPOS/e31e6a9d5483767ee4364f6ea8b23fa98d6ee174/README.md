# STIR/SHAKEN Certificate Repository Compliance

## Rogers Communications Inc

Name: `https://sticr-cstga.ccid.neustar/api/v1/certificate/1da6faf697ce57fff21bed5e9105be2d.crt`\
Tested At: 15 Nov 23 18:07 UTC\
Time: 508ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 15 Nov 23 18:10 UTC