# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Verizon SHAKEN cert 5807

Tested At: 27 Nov 23 22:50 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 81 day(s)\
Subject: CN=Verizon SHAKEN cert 5807, OU=NNO CDS, O=Verizon Data Services LLC, L=Southlake, ST=Texas, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti.verizon.com/vzwcert/vzshaken-02-2024.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICkzCCAjigAwIBAgIQG7m6vUazJ%2Fqvch0J1rR9MTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDIxNjE3MzAxM1oXDTI0MDIxNjE3MzAxM1owgYoxCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVUZXhhczESMBAGA1UEBwwJU291dGhsYWtlMSIwIAYDVQQKDBlWZXJpem9uIERhdGEgU2VydmljZXMgTExDMRAwDgYDVQQLDAdOTk8gQ0RTMSEwHwYDVQQDDBhWZXJpem9uIFNIQUtFTiBjZXJ0IDU4MDcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASBDPoBKI2KeWsVaW9EPaArBrlbiS9%2BAhGAqCDTmkbt520L5IAtToF5rphsKp02yLvpq8Us8Ze9Zh4A%2FYCSZnAho4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ1ODA3MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFLzCDjxfuU3ObZVJBsf27qs8H5TlMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQD1xziWpLrPzBbUyTDWfbafwn7F7YNQSSgcd3l6Yga2BgIhANpoInVtDBD1HC9biSLaQGiQyuBq9ROxyt0qdh9%2FIA7v)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 22:56 UTC