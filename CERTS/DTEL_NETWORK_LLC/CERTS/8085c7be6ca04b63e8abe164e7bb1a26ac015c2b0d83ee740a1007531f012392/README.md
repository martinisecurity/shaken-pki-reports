# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 21 Nov 23 16:41 UTC\
Initial Validity Period: 825 day(s)\
Remaining Validity Period: 755 day(s)\
Subject: CN=SHAKEN, OU=VOIP, O=DTEL NETWORK LLC., L=Sheridan, ST=WY, C=US\
Issuer: emailAddress=info@dtelnetwork.com, O=DTEL NETWORK LLC, L=Default City, C=US, emailAddress=info@dtelnetwork.com\
Link: http://5.78.73.44/system/sp-cert.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIB3jCCAYSgAwIBAgIJAPFLlOnfbhISMAoGCCqGSM49BAMCMGQxCzAJBgNVBAYTAlVTMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxGTAXBgNVBAoMEERURUwgTkVUV09SSyBMTEMxIzAhBgkqhkiG9w0BCQEWFGluZm9AZHRlbG5ldHdvcmsuY29tMB4XDTIzMDkxMjEzNTUxMloXDTI1MTIxNTEzNTUxMlowaTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAldZMREwDwYDVQQHDAhTaGVyaWRhbjEaMBgGA1UECgwRRFRFTCBORVRXT1JLIExMQy4xDTALBgNVBAsMBFZPSVAxDzANBgNVBAMMBlNIQUtFTjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABG6wUA6Aaw0S8TBESGtrDGXbVMmTApyMkiq25nsOOQv%2B9pja5C4t%2FZs8EeSXjlvqdqH2UA43NYU1BbTugMumAxajGjAYMBYGCCsGAQUFBwEaBAowCKAGFgQxMDAxMAoGCCqGSM49BAMCA0gAMEUCIQCpp9T9HHsVb%2B2f%2Bt4Hhpon12E1vKCUzQqtGJ2k2u%2BMIwIgOuTgwxZ1C5WBmYMgdr7rqljumtjBwWoPQpVqeip74v4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1001' |
| [e_atis_ext_authority_key_identifier](../../ISSUES/e_atis_ext_authority_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_subject_key_identifier_size](../../ISSUES/e_atis_subject_key_identifier_size/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_ext_basic_constraints](../../ISSUES/e_atis_ext_basic_constraints/README.md) | error | ATIS1000080 | BasicConstraints extension not found |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_ext_authority_key_identifier_missing](../../ISSUES/e_ext_authority_key_identifier_missing/README.md) | error | RFC5280 |  |
| [e_atis_ext_key_usage](../../ISSUES/e_atis_ext_key_usage/README.md) | error | ATIS1000080 | Key Usage extension not found |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | the Certificate Policies extension is not present |


Generated: 21 Nov 23 17:16 UTC