# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 15 Nov 23 16:04 UTC\
Initial Validity Period: 825 day(s)\
Remaining Validity Period: 761 day(s)\
Subject: CN=SHAKEN, OU=VOIP, O=DTEL NETWORK LLC., L=Sheridan, ST=WY, C=US\
Issuer: emailAddress=info@dtelnetwork.com, O=DTEL NETWORK LLC, L=Default City, C=US, emailAddress=info@dtelnetwork.com\
Link: http://5.78.73.44/system/sp-cert.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIB3jCCAYSgAwIBAgIJAPFLlOnfbhISMAoGCCqGSM49BAMCMGQxCzAJBgNVBAYTAlVTMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxGTAXBgNVBAoMEERURUwgTkVUV09SSyBMTEMxIzAhBgkqhkiG9w0BCQEWFGluZm9AZHRlbG5ldHdvcmsuY29tMB4XDTIzMDkxMjEzNTUxMloXDTI1MTIxNTEzNTUxMlowaTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAldZMREwDwYDVQQHDAhTaGVyaWRhbjEaMBgGA1UECgwRRFRFTCBORVRXT1JLIExMQy4xDTALBgNVBAsMBFZPSVAxDzANBgNVBAMMBlNIQUtFTjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABG6wUA6Aaw0S8TBESGtrDGXbVMmTApyMkiq25nsOOQv%2B9pja5C4t%2FZs8EeSXjlvqdqH2UA43NYU1BbTugMumAxajGjAYMBYGCCsGAQUFBwEaBAowCKAGFgQxMDAxMAoGCCqGSM49BAMCA0gAMEUCIQCpp9T9HHsVb%2B2f%2Bt4Hhpon12E1vKCUzQqtGJ2k2u%2BMIwIgOuTgwxZ1C5WBmYMgdr7rqljumtjBwWoPQpVqeip74v4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain a Subject Key Identifier extension |
| [e_atis_basic_constraints](../../ISSUES/e_atis_basic_constraints/README.md) | error | ATIS1000080 | STI certificates shall contain a BasicConstraints extension marked critical |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | STI certificates shall contain a Key Usage extension marked as critical |
| [e_ext_authority_key_identifier_missing](../../ISSUES/e_ext_authority_key_identifier_missing/README.md) | error | RFC5280 |  |
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1001' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_authority_key_identifier](../../ISSUES/e_atis_authority_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |


Generated: 15 Nov 23 17:17 UTC