# STIR/SHAKEN Certificate Repository Compliance

## EssexTel INC SHAKEN 692J

Name: `https://netnumber-sti-cr.s3.amazonaws.com/certs/419b6f62-c79c-404d-9253-36db8012193f`\
Tested At: 02 Jun 23 01:10 UTC\
Time: 107ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 02 Jun 23 01:12 UTC