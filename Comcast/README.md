# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Leaf Certificates

- Average validity span as of leaf certificates 30 days
- Percentage of leaf certificates expiring in the next 30 days is 100.00%
- Certificates with Errors: 38
- Certificates with Warnings: 38
- Certificates with Notices: 0
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 0

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md#leaf-certificates) | ATIS-1000080 | 38 |
| error | [e_sti_serial_number](ISSUES/e_sti_serial_number/README.md#leaf-certificates) | ATIS-1000080 | 1 |
| error | [e_sti_subject_cn](ISSUES/e_sti_subject_cn/README.md#leaf-certificates) | ATIS-1000080 | 38 |
| error | [e_sti_subject_key_identifier](ISSUES/e_sti_subject_key_identifier/README.md#leaf-certificates) | ATIS-1000080 | 38 |
| warn | [w_cp1_3_subject_rdn_unknown](ISSUES/w_cp1_3_subject_rdn_unknown/README.md#leaf-certificates) | United States SHAKEN CP | 38 |
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md#leaf-certificates) | United States SHAKEN CP | 38 |
| error | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md#leaf-certificates) | United States SHAKEN CP | 38 |
| warn | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md#leaf-certificates) | RFC5280 | 38 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 6 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 30 Sep 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/7e09cbcaca476ca8f980ab8e78118113cb6ef205/README.md) |
| 01 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTIFICATES/0f854cd19526a2ed1f659c3e4f5e2132b0c1ed0e/README.md) |
| 03 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/e45f5da9810dc59bca40951f3b22efc52677dc56/README.md) |
| 03 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTIFICATES/c9c99c097192223c8826a9f3bf5322a0c20a78f4/README.md) |
| 04 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/f73107f6c2178570522cd0246ae1eb2fe882f9ae/README.md) |
| 04 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/fd2814183a001309eed886e8be687c4085fcdc8f/README.md) |
| 05 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/dc9e764717b56faa963a0b026603f04b33542673/README.md) |
| 06 Oct 22 10:05 UTC | SHAKEN | true | [view](CERTIFICATES/303d898eeeb3da9a375ea057a8f1fb1a00921e3f/README.md) |
| 06 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/0228f0f76ce0b557f91c5d005244ed3fd04c8e92/README.md) |
| 06 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/d6ccd2f7b184348e520b673b04912066ce4eab69/README.md) |
| 07 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/c616fc84f5ac07747c5fb094f20cb76f5c146973/README.md) |
| 07 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTIFICATES/e03512878ca79243810bed28c009d73d1b61bdbd/README.md) |
| 08 Oct 22 10:05 UTC | SHAKEN | true | [view](CERTIFICATES/569e81e58638e3e577d9d06a196977333363bf31/README.md) |
| 08 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/87683fc1c307a71a4bbcb85588fc55a7da5f1191/README.md) |
| 09 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/48d021769cdad6cfcb9c754a4c96c173cc0d39e1/README.md) |
| 10 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/e4c9b98bbecb46fde63efe9952112de5b6b99ec1/README.md) |
| 11 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTIFICATES/c204fb1eaedb603ca51eee84b857e46048efef58/README.md) |
| 12 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/051c833d51caccb60906cc06e5008e04c4b9afaa/README.md) |
| 12 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/b88bddd5e19942c872735a02e771f4f996de4b9c/README.md) |
| 13 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/29feec693c178a589259227e49bd7c7080a94722/README.md) |
| 14 Oct 22 10:05 UTC | SHAKEN | true | [view](CERTIFICATES/2a8742b8e2c37f409a0c9c229b8b901248442e37/README.md) |
| 14 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/f51dbe5b93ee71bc386bb2c3e0c6fb2052ca579b/README.md) |
| 14 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTIFICATES/12307fb9d7f33589504a0ac179dd4e66e4c85683/README.md) |
| 15 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/c291effee1f87b2365c2ccdfbf11d3daee6d1793/README.md) |
| 17 Oct 22 10:05 UTC | SHAKEN | true | [view](CERTIFICATES/e3c3476442998a446980b0487de0ecb2da911cf4/README.md) |
| 17 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/93514bcc695cebc98ba110920ceaea4ee7fc5437/README.md) |
| 17 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/caa3abae9e9defc436e4c1208744ec71587c5e1b/README.md) |
| 18 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTIFICATES/0a21bcc994571022166cd567493f888abadf1bac/README.md) |
| 18 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTIFICATES/ad247d39f3025524d72756ee747c72aa9d2efcab/README.md) |
| 19 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/b11ce2599d4b279750f8823c8f39c6e4540afafe/README.md) |
| 20 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTIFICATES/f9fdf7a4cb279213b0cb0d2f0638caa19dd4c36c/README.md) |
| 20 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/8e62e5ab2267e22f4d55a4ea43494af79ce80900/README.md) |
| 21 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTIFICATES/78c0ceb1e0b0d78a459ddd4bd46c08017b9a07f4/README.md) |
| 21 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTIFICATES/7adb08c0422dfdbc535f11c8f108e6110df249ea/README.md) |
| 22 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTIFICATES/0e0b2467e9112c6a396bbe66266203b592db93aa/README.md) |
| 24 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTIFICATES/02a8317e417bc66baac5ec69fcc42e5a3613a834/README.md) |
| 25 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTIFICATES/3407ac0a1a3530ed984de90a1518de2b8bc1b13a/README.md) |
| 26 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTIFICATES/b9a2a1554c1e26638a56b319ee1aa6b2837fafbe/README.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL/README.md).

### CA Certificates

- Certificates with Errors: 0
- Certificates with Warnings: 0
- Certificates with Notices: 1
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 2

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_ca_subject_public_key](ISSUES/e_sti_ca_subject_public_key/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_authority_key_identifier](ISSUES/e_sti_ca_authority_key_identifier/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_version](ISSUES/e_sti_ca_version/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_issuer_dn](ISSUES/e_sti_ca_issuer_dn/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_key_usage](ISSUES/e_sti_ca_key_usage/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_root_authority_key_identifier](ISSUES/e_sti_root_authority_key_identifier/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_subject](ISSUES/e_sti_ca_subject/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md#ca-certificates) | ATIS-1000080 | 2 |
| notice | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md#ca-certificates) | SHAKEN PKI Best Practice | 1 |
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md#ca-certificates) | United States SHAKEN CP | 2 |
| not effective | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_signature_algorithm](ISSUES/e_sti_ca_signature_algorithm/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [w_cp1_3_ca_subject_rdn_unknown](ISSUES/w_cp1_3_ca_subject_rdn_unknown/README.md#ca-certificates) | United States SHAKEN CP | 2 |
| not effective | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_subject_key_identifier](ISSUES/e_sti_ca_subject_key_identifier/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 1 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 17 Mar 20 14:45 UTC | Comcast SHAKEN Root CA | false | [view](CERTIFICATES/e341fff079ef701a75085e21aaa915d84a27a52a/README.md) |
| 06 Apr 20 13:48 UTC | Comcast SHAKEN Intermediate CA | false | [view](CERTIFICATES/39011602b92be825bea3e29648f2e1866d60d0c6/README.md) |

Generated: 27/10/2022 at 21:27:34