# Digital-Signature-Sample
Digital Signature Based On Bitcoin Address with Golang

목표
비트코인의 주소 체계를 기반으로 임의의 데이터의 전자서명 검증 로직 개발


내용
1. ECDSA기반 key pair (private, public key) 생성
2. 비트코인의 주소 체계와 같은 주소 생성
3. 임의의 데이터(문자열)을 private key로 signature 생성
4. Public key를 이용하여 검증
