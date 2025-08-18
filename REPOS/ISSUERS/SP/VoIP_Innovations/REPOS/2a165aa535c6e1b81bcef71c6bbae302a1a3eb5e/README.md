# STIR/SHAKEN Certificate Repository Compliance

## VoIP Innovations

Name: `https://ssc.getsipnav.com/certs/df5936cc4f0da18288b02b726643bd9c728f6c51`\
Tested At: 18 Aug 25 21:06 UTC\
Time: 568ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC