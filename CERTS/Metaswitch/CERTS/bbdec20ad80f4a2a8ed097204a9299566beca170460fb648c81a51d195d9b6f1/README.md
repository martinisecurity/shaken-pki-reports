# STIR/SHAKEN CA Ecosystem Compliance

## Certificate RCN SHAKEN Cert 7615

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: -201 day(s)\
Subject: CN=RCN SHAKEN Cert 7615, O=RCN, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/91cb0b66fe231e4eb06ddad650a80f4f09cdfdc4

[View certificate details](https://x509.io/?cert=MIICQTCCAeegAwIBAgIQCFfTK6Nnq8%2BDNtDYu%2BZYnzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxNzE3MDYwMVoXDTI0MDMxNjE3MDYwMVowOjELMAkGA1UEBhMCVVMxDDAKBgNVBAoMA1JDTjEdMBsGA1UEAwwUUkNOIFNIQUtFTiBDZXJ0IDc2MTUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATAt6ELjr2XfUM%2FjKVUUjVMaZcM45EZ%2FuxWR4jx0mAVh8Q6Uvfrxq8h8gnbLVtqhV%2FRMz0VNdsrqrG0WUpG4X3po4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3NjE1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHd6jgfsvzke%2Bwfzg1VrzP25KaPeMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCICUcEjuJCsxEFC6zuTL7lSoVrTdYAyCDB%2Be8qnO4bol1AiEAtX1wCQcGbJmq4L9Hnu2WmV5xO1m4ccoKQHdKu52oYh0%3D)

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


Generated: 04 Oct 24 16:29 UTC