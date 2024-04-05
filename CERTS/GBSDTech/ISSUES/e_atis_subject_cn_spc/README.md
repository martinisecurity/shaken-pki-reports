# STIR/SHAKEN CA Ecosystem Compliance

## GBSDTech

Name: e_atis_subject_cn_spc\
Source: ATIS1000080\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 16 Jan 22 00:00 UTC\
Description: For end-entity certificate, the Common Name attribute shall contain the text string SHAKEN, followed by a single space, followed by the SPC value identified in the TNAuthList of the end-entity certificate.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | FracTEL LLC SHAKEN | [view](../../CERTS/7e9f3377c97ba96475fb319a39341fb126f4d96175b09c0ea388618ee1749e50/README.md) | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| error | Edify SHAKEN | [view](../../CERTS/d092ee80d10eb8c6656246f9ffa3d2100319fb217c50c8cc03e4d84e654da026/README.md) | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| error | ccctelecom.com | [view](../../CERTS/e485e7069de2a669c6cd53d3eb10263d437b1c616c03e7e2997a531d020361ad/README.md) | Common name shall contain the text string 'SHAKEN 816K', but common name is 'ccctelecom.com' |
| error | MYPBXManager SHAKEN | [view](../../CERTS/ea5813855308274fae05fdcae622a159efa47cde2ccf87a9cdf09d9ef43d93f2/README.md) | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 05 Apr 24 19:04 UTC