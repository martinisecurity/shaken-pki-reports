# STIR/SHAKEN Certificate Repository Compliance

## AM Communications Labs Inc

Name: `https://ssc.getsipnav.com/certs/d68017fb427d47359b7108377a18cf105aad7709`\
Tested At: 21 Nov 23 01:49 UTC\
Time: 101ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Nov 23 01:55 UTC