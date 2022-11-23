# STIR/SHAKEN Certificate Repository Compliance

## Pulsar360 Corp

Name: `https://ecms.securetransit.net/certs/410bad9e05a57353f0d8739a421be619.cer`\
Tested At: 23 Nov 22 18:08 UTC\
Time: 373ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 23 Nov 22 18:09 UTC