# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Leaf Certificates

- Average validity span as of leaf certificates 82 days
- Percentage of leaf certificates expiring in the next 30 days is 74.60%
- Certificates with Errors: 63
- Certificates with Warnings: 62
- Certificates with Notices: 0
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 2

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn.md#leaf-certificates) | CPv1.3 | 63 |
| error | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier.md#leaf-certificates) | CPv1.3 | 4 |
| not effective | [e_sti_key_usage](ISSUES/e_sti_key_usage.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_authority_key_identifier](ISSUES/e_sti_authority_key_identifier.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| not effective | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email.md#leaf-certificates) | CPv1.3 | 1 |
| warn | [w_cp1_3_subject_rdn_unknown](ISSUES/w_cp1_3_subject_rdn_unknown.md#leaf-certificates) | CPv1.3 | 63 |
| not effective | [e_sti_signature_algorithm](ISSUES/e_sti_signature_algorithm.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_serial_number](ISSUES/e_sti_serial_number.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_extension_unknown](ISSUES/e_sti_extension_unknown.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_subject_key_identifier](ISSUES/e_sti_subject_key_identifier.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| not effective | [n_sti_certificate_policy_critical](ISSUES/n_sti_certificate_policy_critical.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies.md#leaf-certificates) | ATIS-1000080v4 | 4 |
| not effective | [e_sti_version](ISSUES/e_sti_version.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_subject_public_key](ISSUES/e_sti_subject_public_key.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_subject_cn](ISSUES/e_sti_subject_cn.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_issuer_dn](ISSUES/e_sti_issuer_dn.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| error | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding.md#leaf-certificates) | RFC5280 | 63 |
| error | [e_sti_crl_distribution](ISSUES/e_sti_crl_distribution.md#leaf-certificates) | ATIS-1000080v4 | 63 |
| not effective | [e_sti_tn_auth_list](ISSUES/e_sti_tn_auth_list.md#leaf-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_subject](ISSUES/e_sti_subject.md#leaf-certificates) | ATIS-1000080v4 | 2 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 229 certificates skipped because they are currently expired.\
\*\*\* 1 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 02 Nov 21 16:40 UTC | Charter Communications Inc SHAKEN 5606 | true | [view](abbd6bca303f79a930d3a14e0d3538c51934a97a%2FREADME.md) |
| 13 Jan 22 14:17 UTC | MobileSphere SHAKEN 873J | true | [view](701a8780085449f2f6ee50c287cecc20f2219b2f%2FREADME.md) |
| 30 Mar 22 12:54 UTC | Fusion Connect SHAKEN 2720 | true | [view](136f6839b2fa440965940ff380d9cde7b053f2d0%2FREADME.md) |
| 27 Apr 22 18:22 UTC | CCI SHAKEN 663J | true | [view](7c0be1458e20b8247bce9338ecbe8322578f807c%2FREADME.md) |
| 10 Jun 22 14:00 UTC | SHAKEN 597J | true | [view](9e8c6e83efe9133df13009a69d1dfbe8801f0a13%2FREADME.md) |
| 20 Jun 22 20:25 UTC | SHAKEN 857J | true | [view](caf6b73e79063f79bf34dae854ebf059f32c1a45%2FREADME.md) |
| 22 Jun 22 18:46 UTC | SHAKEN 755J | true | [view](ee364dffdc58cce8eb91842e65530ece378acec1%2FREADME.md) |
| 29 Jun 22 20:24 UTC | SHAKEN 736J | true | [view](8a1091afcbec190e8877ecdd666ef18c50cb9356%2FREADME.md) |
| 25 Jul 22 18:36 UTC | SHAKEN 578J | true | [view](b0affbea73591da4097b4915d92b53818409f67f%2FREADME.md) |
| 10 Aug 22 18:11 UTC | SHAKEN 073H | true | [view](0ae828173955498af37d944351d38f36446db6d6%2FREADME.md) |
| 01 Sep 22 03:25 UTC | SHAKEN 6628 | true | [view](01505cd524c060dffca87d0b043e0be889026713%2FREADME.md) |
| 01 Sep 22 12:31 UTC | SHAKEN 193E | true | [view](b6d72ada27949e72427f323a03e4ddc7e6d2e803%2FREADME.md) |
| 13 Sep 22 17:24 UTC | SHAKEN 706J | true | [view](8c58c1679cf840f228b7d5285f431a25ad4d6562%2FREADME.md) |
| 26 Sep 22 18:40 UTC | SHAKEN 505J | true | [view](b110a3f4054ff7065b1c6f14883ade56f2ced44e%2FREADME.md) |
| 26 Sep 22 18:42 UTC | SHAKEN 505J | true | [view](d3e15676cd711e2afcfebdfa1c0ea2186d50b5b9%2FREADME.md) |
| 30 Sep 22 22:34 UTC | SHAKEN 8526 | true | [view](969863c8a4f335ed0a19906ae5c3558905b7725a%2FREADME.md) |
| 01 Oct 22 04:04 UTC | SHAKEN 345J | true | [view](eb03b74587d3cfb43d588d8463e17c52db9244ed%2FREADME.md) |
| 01 Oct 22 04:34 UTC | SHAKEN 345J | true | [view](22a8d66576c72f5acfc282c00a7f9e28863234ea%2FREADME.md) |
| 01 Oct 22 12:31 UTC | SHAKEN 193E | true | [view](164675f5894b5459867b897576a59e6aba1fc656%2FREADME.md) |
| 01 Oct 22 15:23 UTC | SHAKEN 722J | true | [view](006cd8457db666bfa0b1df0497d439dd83f9f246%2FREADME.md) |
| 10 Oct 22 18:30 UTC | SHAKEN 115K | true | [view](dd733ba189741f0b583a3e595a3c14db8e87be01%2FREADME.md) |
| 12 Oct 22 19:39 UTC | SHAKEN 815G | true | [view](e689269bc057ddf1876d610f8dd14761b439defc%2FREADME.md) |
| 13 Oct 22 18:02 UTC | SHAKEN 952J | true | [view](414512f13c27bb40a6f47741eabd63eb5c90eed3%2FREADME.md) |
| 20 Oct 22 09:13 UTC | SHAKEN 841J | true | [view](3889bee709733ff19f8acc2ae4a3761d288ffdfe%2FREADME.md) |
| 20 Oct 22 15:48 UTC | SHAKEN 622J | true | [view](175e3b3de9a38a8355d047b901bd5eda0c1f1a6a%2FREADME.md) |
| 20 Oct 22 18:13 UTC | SHAKEN 060K | true | [view](0b14a0d19d2f23312c25f7a36b56acb673cc9762%2FREADME.md) |
| 20 Oct 22 20:17 UTC | SHAKEN 674J | true | [view](63a29d037569770d076f719fd55e4daa34c1b717%2FREADME.md) |
| 20 Oct 22 20:18 UTC | SHAKEN 738J | true | [view](ffe9e32959b31e1ef27b20255df21d5408b7c432%2FREADME.md) |
| 20 Oct 22 20:21 UTC | SHAKEN 854D | true | [view](54c7bf84f9870adf34d3624c71a4b8224436f994%2FREADME.md) |
| 20 Oct 22 20:21 UTC | SHAKEN 790J | true | [view](992671c6facefac133d44aa22b9d4f09f46456e8%2FREADME.md) |
| 20 Oct 22 20:22 UTC | SHAKEN 459J | true | [view](8c970d3ca70f330e52327f79d8039b381d94826f%2FREADME.md) |
| 20 Oct 22 20:23 UTC | SHAKEN 366G | true | [view](b4d26c19a9c0aef72e9326e6bcf76f925cc71b01%2FREADME.md) |
| 20 Oct 22 20:23 UTC | SHAKEN 738J | true | [view](b394a8c34e7c0ca1568a3be4c04c8399e79156d9%2FREADME.md) |
| 20 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](8a595b01c0d9c2558e1a00b65da85badc3d941d2%2FREADME.md) |
| 21 Oct 22 20:13 UTC | SHAKEN 983J | true | [view](fa5b20f5d1c96d0095d20cd3cf723e84a876e7ab%2FREADME.md) |
| 22 Oct 22 00:21 UTC | SHAKEN 551G | true | [view](fc9938ac7cc123cd2e27f38171de9b31183c5871%2FREADME.md) |
| 22 Oct 22 04:04 UTC | SHAKEN 345J | true | [view](fc99f74650aae6ff90f41e781e215bf1a8f06f68%2FREADME.md) |
| 22 Oct 22 04:44 UTC | SHAKEN 345J | true | [view](ecf5e6662e3084505582f48a356cc82f2e95a7dd%2FREADME.md) |
| 23 Oct 22 08:39 UTC | SHAKEN 012K | true | [view](a96d0880d34d1cb60fd0f00ce985722fb952a02d%2FREADME.md) |
| 23 Oct 22 18:35 UTC | SHAKEN 089K | true | [view](fde031a4737afd05ee9408293b200a35bfa307e4%2FREADME.md) |
| 23 Oct 22 18:37 UTC | SHAKEN 9714 | true | [view](cb1430446ee425f297b477264f245934a792e826%2FREADME.md) |
| 23 Oct 22 19:56 UTC | SHAKEN 297K | true | [view](9df80eed66106e1aa17a18c41d830a3e240c445a%2FREADME.md) |
| 23 Oct 22 20:17 UTC | SHAKEN 674J | true | [view](ff10d41f263188f0218fa8d9f25b21d42e3fe1f7%2FREADME.md) |
| 23 Oct 22 20:18 UTC | SHAKEN 738J | true | [view](bf14dbda90b2ffc796c34770c105c1ca625de56b%2FREADME.md) |
| 23 Oct 22 20:18 UTC | SHAKEN 7453 | true | [view](e6d88e273d127842c50b03c4b5e1c0dd088872fe%2FREADME.md) |
| 23 Oct 22 20:21 UTC | SHAKEN 849J | true | [view](ef4f17c2145aaea571f41b690925b622e8146800%2FREADME.md) |
| 23 Oct 22 20:21 UTC | SHAKEN 469A | true | [view](b3972d4d4686a49e9ae279e328f8b822c476fe52%2FREADME.md) |
| 23 Oct 22 20:21 UTC | SHAKEN 790J | true | [view](38fe760cf3e12683c131b3d0f2e891f007794503%2FREADME.md) |
| 23 Oct 22 20:22 UTC | SHAKEN 625J | true | [view](e2aacd36d02951b3e73b552ea16b976f7ccf9f26%2FREADME.md) |
| 23 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](dcfb7c5055b8e4f07f3136116eaf7e354dabb82a%2FREADME.md) |
| 23 Oct 22 20:23 UTC | SHAKEN 459J | true | [view](cf6536ed84fa4fb421ebf66f89eb4d51794270c3%2FREADME.md) |
| 23 Oct 22 20:24 UTC | SHAKEN 366G | true | [view](bd7a8b9542e2a04e92af34f261cb8afa92a1af0d%2FREADME.md) |
| 23 Oct 22 20:24 UTC | SHAKEN 0226 | true | [view](f9c467c1018ad1eabfa846fc2aa1c5c27cd8123f%2FREADME.md) |
| 23 Oct 22 20:24 UTC | SHAKEN 738J | true | [view](e8f019de24ae5f9dd28301a9f73bb0ce59907a7f%2FREADME.md) |
| 23 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](76579174555c2405b844fbf6ace0ccb15e0aa02c%2FREADME.md) |
| 23 Oct 22 22:11 UTC | SHAKEN 0172 | true | [view](8dcc3b92f9308b0c51305675d84fe107f0187ecc%2FREADME.md) |
| 24 Oct 22 16:32 UTC | SHAKEN 291K | true | [view](2661ebd23f777002d5a8b8e0cced45b06bf12b10%2FREADME.md) |
| 24 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](aa9b65dead966d217c2e51f1e78559a830fefbf8%2FREADME.md) |
| 24 Oct 22 20:51 UTC | SHAKEN 157C | true | [view](21e6797a2a553ddf3e1d5c9d5ceb99f93fa6c773%2FREADME.md) |
| 24 Oct 22 21:41 UTC | SHAKEN 606F | true | [view](0ea62e7be98fcc1d31d5d28965f8fabeb1cd9e5a%2FREADME.md) |
| 25 Oct 22 15:33 UTC | SHAKEN 722J | true | [view](160aff282003da74ec5b92fd7f622fb8487e6879%2FREADME.md) |
| 26 Oct 22 08:39 UTC | SHAKEN 012K | true | [view](d645bcf28ab9c9c17f50a509aee493520976cd68%2FREADME.md) |
| 26 Oct 22 10:43 UTC | SHAKEN 8526 | true | [view](c75227976fe6d09fed2d445131b0aa220b68e55f%2FREADME.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL.md).

### CA Certificates

- Certificates with Errors: 3
- Certificates with Warnings: 0
- Certificates with Notices: 3
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 3

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints.md#ca-certificates) | ATIS-1000080v4 | 3 |
| not effective | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_authority_key_identifier](ISSUES/e_sti_ca_authority_key_identifier.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn.md#ca-certificates) | ATIS-1000080v4 | 3 |
| not effective | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number.md#ca-certificates) | ATIS-1000080v4 | 3 |
| not effective | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical.md#ca-certificates) | ATIS-1000080v4 | 2 |
| error | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding.md#ca-certificates) | RFC5280 | 3 |
| not effective | [w_cp1_3_ca_subject_rdn_unknown](ISSUES/w_cp1_3_ca_subject_rdn_unknown.md#ca-certificates) | CPv1.3 | 3 |
| not effective | [e_sti_ca_subject](ISSUES/e_sti_ca_subject.md#ca-certificates) | ATIS-1000080v4 | 3 |
| not effective | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_signature_algorithm](ISSUES/e_sti_ca_signature_algorithm.md#ca-certificates) | ATIS-1000080v4 | 3 |
| not effective | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_version](ISSUES/e_sti_ca_version.md#ca-certificates) | ATIS-1000080v4 | 3 |
| notice | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage.md#ca-certificates) | SHAKEN PKI Best Practice | 3 |
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign.md#ca-certificates) | CPv1.3 | 3 |
| not effective | [e_sti_ca_issuer_dn](ISSUES/e_sti_ca_issuer_dn.md#ca-certificates) | ATIS-1000080v4 | 3 |
| not effective | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_subject_public_key](ISSUES/e_sti_ca_subject_public_key.md#ca-certificates) | ATIS-1000080v4 | 3 |
| not effective | [e_sti_ca_key_usage](ISSUES/e_sti_ca_key_usage.md#ca-certificates) | ATIS-1000080v4 | 3 |
| not effective | [e_sti_ca_subject_key_identifier](ISSUES/e_sti_ca_subject_key_identifier.md#ca-certificates) | ATIS-1000080v4 | 3 |
| not effective | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown.md#ca-certificates) | ATIS-1000080v4 | 2 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 1 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Root CA1 | true | [view](36dc4ae1d521b8a5aedd10498e6ce757581b197f%2FREADME.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA1 | true | [view](2d540c7179aad19290a2098ccac8053645527154%2FREADME.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA3 | true | [view](df7d871ff60d213820a96308346eda870d6e8ed2%2FREADME.md) |

Generated: 26/10/2022 at 22:31:35