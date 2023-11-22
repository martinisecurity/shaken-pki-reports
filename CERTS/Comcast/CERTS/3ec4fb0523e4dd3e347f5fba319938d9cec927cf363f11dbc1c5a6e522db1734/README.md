# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 22 Nov 23 03:34 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -93 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/3a4407fb-c4c3-400b-aab8-58a54635c656.cer

[View certificate details](https://understandingwebpki.com/?cert=MIICWDCCAf2gAwIBAgIIRabjyvzK4iQwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIzMDcyMTEzNTUwMFoXDTIzMDgyMDEzNTUwMFowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARbFmSx5Pmj%2Bx4NmdDUCBf9jxrZOx%2B52VKjZuk9P2UZEzldFPkVQmNGSEnKKCZCuqbFh%2BBnfSvshCvYyxuiOrSOo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYEMzE4SjAKBggqhkjOPQQDAgNJADBGAiEAw%2BQkoP1pkgoA8rRfkjW36EBo6a%2FMkjUpt2oFGTjU9tECIQDlrgHIfiPEUThwz9Ilt5rtGm94CYBEPxRPfW%2FIlY%2Bfpg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 318J' |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | the Certificate Policies extension is not present |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_subject_key_identifier_size](../../ISSUES/e_atis_subject_key_identifier_size/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |


Generated: 22 Nov 23 03:38 UTC