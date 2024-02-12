# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ENA SHAKEN cert 521F

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 123 day(s)\
Subject: CN=ENA SHAKEN cert 521F, O=ENA, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/702fb83e8268b5175b591d8448a512646339bc40

[View certificate details](https://understandingwebpki.com/?cert=MIICQTCCAeegAwIBAgIQH2O8LTasSJz0F8wWympA3TAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYxNTEzNTY1MVoXDTI0MDYxNDEzNTY1MVowOjELMAkGA1UEBhMCVVMxDDAKBgNVBAoMA0VOQTEdMBsGA1UEAwwURU5BIFNIQUtFTiBjZXJ0IDUyMUYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATww3ygb81IGEdAyq4Dc0OJEUAOF3kLhGB9HJF0xOihIvMmHQqpO8v95IEi0tt8usB7%2Bx2w0FadPsBYjzPZ%2F6ylo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ1MjFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHkOd0WGEQe5gWecn95SH%2FvhOqbnMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQDmdX3vvawVcYCD0dgkIjqmKcaa2wD%2FUfNIEe31fAWtCwIgaYXGNP25PAhY9k%2F3tlyHAtTlETsEqViAVFECul%2FuuLE%3D)

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


Generated: 12 Feb 24 17:02 UTC