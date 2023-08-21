# STIR/SHAKEN Certificate Repository Compliance

## Accessline Communications Corp.

Name: `https://sscerts.accessline.com/sscerts/Cert-US-04-2022.crt`\
Tested At: 21 Aug 23 20:16 UTC\
Time: 188ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Aug 23 20:18 UTC