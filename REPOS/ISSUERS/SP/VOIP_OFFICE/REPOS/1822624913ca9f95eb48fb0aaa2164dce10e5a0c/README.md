# STIR/SHAKEN Certificate Repository Compliance

## VOIP OFFICE

Name: `https://nguc.voipoffice.com/stirshaken/VOIP_OFFICE.COM_LLC_389K`\
Tested At: 27 Nov 23 23:21 UTC\
Time: 49ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 27 Nov 23 23:28 UTC