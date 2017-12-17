import time
from datasketch.minhash import MinHash

def main():
    with open('plato1.txt', 'r') as f:
        tokens1 = [l for l in f]
    with open('plato2.txt', 'r') as f:
        tokens2 = [l for l in f]

    start = time.time()
    m1 = MinHash(num_perm=64, seed=0)
    for t in tokens1:
        m1.update(t.encode('utf8'))

    m2 = MinHash(num_perm=64, seed=0, permutations=m1.permutations)
    for t in tokens2:
        m2.update(t.encode('utf8'))
    similarity = m2.jaccard(m1)
    elapsed = time.time() - start
    print("Similar %f and Took %f ms", similarity, elapsed*1000)

if __name__ == "__main__":
    main()
