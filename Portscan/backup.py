Usar quibits em vez de bits para realizar um scan 
de portas não deve confundir a detecção de firewall
de forma significativa. A principal diferença 
entre o uso de bits clássicos e quibits quânticos é
a maneira como as informações são processadas e representadas.

Ao realizar um scan de portas usando quibits, estamos
aproveitando os princípios da computação quântica para 
realizar cálculos de forma paralela e explorar propriedades 
como emaranhamento e superposição.

Emaranhamento (Entanglement): O emaranhamento refere-se a
uma propriedade quântica na qual dois ou mais qubits
(quanto bits) tornam-se intrinsecamente correlacionados, 
independentemente da distância entre eles. Isso significa 
que as mudanças em um qubit afetarão instantaneamente o estado
do outro, mesmo que estejam separados por grandes distâncias.
O emaranhamento permite que a informação seja codificada de 
maneira complexa e compartilhada entre os qubits, fornecendo 
uma vantagem computacional significativa em certos
tipos de cálculos quânticos.

Superposição (Superposition): A superposição é outra
propriedade quântica que permite que um qubit esteja em um
estado de combinação linear de diferentes estados simultaneamente
Enquanto os bits clássicos podem estar em um estado 0 ou 1, 
um qubit em superposição pode estar em uma combinação de 0 e 1
ao mesmo tempo. Isso significa que um qubit pode existir em uma
sobreposição de vários estados possíveis, aumentando 
significativamente a capacidade de processamento e armazenamento 
de informações.

b'\x18U\x08\xac\x07\xf7\x01?\x1bw2P\x00)\x05r\xaaa\x99py\x7f\x9f\xac\xd8\xfe\x07\xfc1\x12Q-'

# AtlasProject - Quantum Portscan AES-256 Payload


from qiskit import QuantumCircuit, execute, Aer
from scapy.all import IP, send

def scan_quantico(ip_destino):
    num_qubits = 4

    circuito = QuantumCircuit(num_qubits, num_qubits)

    for i in range(num_qubits):
        circuito.h(i)

    for i in range(num_qubits - 1):
        circuito.cx(i, num_qubits - 1)

    circuito.measure(range(num_qubits), range(num_qubits))

    simulador = Aer.get_backend('qasm_simulator')
    job = execute(circuito, simulador, shots=1)
    resultado = job.result()

    contagens = resultado.get_counts(circuito)

    portas_abertas = []
    for estado, contagem in contagens.items():
        if estado[num_qubits - 1] == '1':
            porta = int(estado[:num_qubits - 1], 2)
            portas_abertas.append(porta)

    if portas_abertas:
        print("Portas abertas:")
        for porta in portas_abertas:
            if porta == 0:
                print("Porta", porta, "aberta (22).")
            elif porta == 1:
                print("Porta", porta, "aberta (80).")
            elif porta == 2:
                print("Porta", porta, "aberta (443).")
            elif porta == 3:
                print("Porta", porta, "aberta (3389).")
            elif porta == 4:
                print("Porta", porta, "aberta (8080).")
            elif porta == 5:
                print("Porta", porta, "aberta (8000).")
            elif porta == 6:
                print("Porta", porta, "aberta (21).")
            elif porta == 7:
                print("Porta", porta, "aberta (3306).")
    else:
        print("Nenhuma porta aberta encontrada.")

def enviar_pacote_disfarcado(ip_destino):
    pacote_disfarcado = IP(dst=ip_destino) / "b'\x18U\x08\xac\x07\xf7\x01?\x1bw2P\x00)\x05r\xaaa\x99py\x7f\x9f\xac\xd8\xfe\x07\xfc1\x12Q-'"

    send(pacote_disfarcado)

    print("Pacote disfarçado enviado com sucesso para", ip_destino)


ip_destino = "200.1.183.71"
scan_quantico(ip_destino)
enviar_pacote_disfarcado(ip_destino)







