# STIR/SHAKEN CA Ecosystem Compliance

## Certificate 1RouteGroup SHAKEN Intermediate CA

Tested At: 26 Aug 24 18:48 UTC\
Initial Validity Period: 7299 day(s)\
Remaining Validity Period: 6091 day(s)\
Subject: CN=1RouteGroup SHAKEN Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Issuer: CN=GBSDTech SHAKEN Root CA, O=GBSDTech, L=Ft Worth, ST=Texas, C=US

[View certificate details](https://x509.io/?cert=MIIB9zCCAZ2gAwIBAgICEAAwCgYIKoZIzj0EAwIwZTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMREwDwYDVQQHDAhGdCBXb3J0aDERMA8GA1UECgwIR0JTRFRlY2gxIDAeBgNVBAMMF0dCU0RUZWNoIFNIQUtFTiBSb290IENBMB4XDTIxMDUwNTIwMjIyNVoXDTQxMDQyOTIwMjIyNVowYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiMVJvdXRlR3JvdXAgU0hBS0VOIEludGVybWVkaWF0ZSBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNKPuHvnjJ6FNM5o%2Bjj781vtx0UAotBJJjAiSXslRwm8oTxEshFnChBJ%2FwM6eBVoQ6eL6mcQb7ZgoZ7k%2BNi%2BQ6KjQjBAMB0GA1UdDgQWBBRFCSVsXS7H4VCeIIYTuSTGc0%2FiPjAPBgNVHRMBAf8EBTADAQH%2FMA4GA1UdDwEB%2FwQEAwICBDAKBggqhkjOPQQDAgNIADBFAiAWnq2bNOqoldt%2Bk9te%2Bo2UwAP2gP8KtxWFhHjsnxDxoAIhAK2oWP4XjNs%2FsYmDNNiOdwFAO%2B9t83efVBYCY3tUYHlS)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_ca](../../ISSUES/e_atis_ext_crl_distribution_ca/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |
| [e_atis_ext_authority_key_identifier_ca](../../ISSUES/e_atis_ext_authority_key_identifier_ca/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |
| [e_ext_authority_key_identifier_missing](../../ISSUES/e_ext_authority_key_identifier_missing/README.md) | error | RFC5280 |  |
| [e_atis_ext_certificate_policies_ca](../../ISSUES/e_atis_ext_certificate_policies_ca/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |

### Not Effective

- e_atis_ext_not_specified_ca
- e_atis_serial_number_size_ca
- e_atis_subject_c_iso_ca
- e_atis_subject_key_identifier_size_ca
- e_atis_subject_o_required_ca
- e_shaken_certificate_policies_id_ca

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:49 UTC