from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import padding
import os

def encrypt_AES(message, key):
    padder = padding.PKCS7(128).padder()
    padded_data = padder.update(message.encode()) + padder.finalize()

    iv = os.urandom(16)

    cipher = Cipher(algorithms.AES(key), modes.CBC(iv), backend=default_backend())

    encryptor = cipher.encryptor()

    ciphertext = encryptor.update(padded_data) + encryptor.finalize()

    return iv + ciphertext
chave = os.urandom(32)

# Mensagem a ser criptografada
mensagem = "Olá, mundo"

# Criptografa a mensagem usando AES-256
texto_cifrado = encrypt_AES(mensagem, chave)

print("Texto cifrado (AES-256):", texto_cifrado)
