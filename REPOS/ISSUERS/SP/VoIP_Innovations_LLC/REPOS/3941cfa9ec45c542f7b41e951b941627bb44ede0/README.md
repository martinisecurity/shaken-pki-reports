# STIR/SHAKEN Certificate Repository Compliance

## VoIP Innovations LLC

Name: `https://ssc.getsipnav.com/certs/4271238eaebdb4a770896e8a3a4d18618f648cb6`\
Tested At: 04 Oct 24 16:23 UTC\
Time: 142ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC