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
| error | NovoLink SHAKEN cert | [view](../../CERTS/2a55030811ea276fc007055026d66293c349f35c5eebc7dd7b305b65f3b40c56/README.md) | Common name shall contain the text string 'SHAKEN 332G', but common name is 'NovoLink SHAKEN cert' |
| error | FracTEL LLC SHAKEN | [view](../../CERTS/436ab7f1876a96dfebbc3cdb0880c6f4dc9432a1db1bec8745e39d13d4febc22/README.md) | Common name shall contain the text string 'SHAKEN 965H', but common name is 'FracTEL LLC SHAKEN' |
| error | alluretelecom.com | [view](../../CERTS/5486360ac5b339b588547ee88d42ca59ee3cdd87a0911ce2d16c3a8dd376efe0/README.md) | Common name shall contain the text string 'SHAKEN 846K', but common name is 'alluretelecom.com' |
| error | ccctelecom.com | [view](../../CERTS/e485e7069de2a669c6cd53d3eb10263d437b1c616c03e7e2997a531d020361ad/README.md) | Common name shall contain the text string 'SHAKEN 816K', but common name is 'ccctelecom.com' |
| error | MYPBXManager SHAKEN | [view](../../CERTS/ea5813855308274fae05fdcae622a159efa47cde2ccf87a9cdf09d9ef43d93f2/README.md) | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 26 Aug 24 18:49 UTC