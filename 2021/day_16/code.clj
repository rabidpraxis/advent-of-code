(ns day-16.code
  (:require
    [clojure.string :as str]))

(def final-data
  (slurp "day_16/input.txt"))

(def bits
  {"0" "0000"
   "1" "0001"
   "2" "0010"
   "3" "0011"
   "4" "0100"
   "5" "0101"
   "6" "0110"
   "7" "0111"
   "8" "1000"
   "9" "1001"
   "A" "1010"
   "B" "1011"
   "C" "1100"
   "D" "1101"
   "E" "1110"
   "F" "1111" })

(defn parse-binary [s]
  (Long/parseLong s 2))

(defn s-rest [s]
  (subs s 1 (count s)))

(defn extract [s len]
  [(subs s 0 len) (subs s len (count s))])

(defn parse-hex [s]
  (->> (str/split s #"") (map bits) (str/join "" )))

(defn flatten-ops [s]
  (lazy-seq
    (if (:ops s)
      (cons (dissoc s :ops) (mapcat flatten-ops (:ops s)))
      (list s))))

(defn packet [s]
  (let [[v s]   (extract s 3)
        version (parse-binary v)
        [t s]   (extract s 3)
        tval    (parse-binary t)
        type    (if (= tval 4) :literal :operator)]
    {:version version
     :type-val tval
     :type type
     :rest s }))

(defn literal-process [{:keys [rest] :as p}]
  (loop
    [s rest
     len 6
     final []]
    (let [[chunk s] (extract s 5)
          final (conj final (s-rest chunk)) ]
      (if (= (first chunk) \0)
        (assoc p
               :rest s
               :len (+ len 5)
               :value (parse-binary (str/join "" final))
               )
        (recur s (+ len 5) final)))))

(defn operator-process [{:keys [rest] :as p}]
  (let [[v s] (extract rest 1)]
    (if (= v "0")
      (let [[n s] (extract s 15)]
        (assoc p
               :op-type :length
               :len (parse-binary n)
               :rest s))
      (let [[n s] (extract s 11)]
        (assoc p
               :op-type :count
               :ct (parse-binary n)
               :rest s)))))

(defmulti process-op :op-type)
(defmulti process :type)
(defmulti calc-value :type-val)

(defmethod process-op :count [{:keys [rest ct] :as s}]
  (let [ops (loop [ct ct rest rest final []]
              (let [processed (process (packet rest))]
                (if (zero? (dec ct))
                  (conj final processed)
                  (recur (dec ct) (:rest processed) (conj final processed)))))
        nlen (apply + (map :len ops))]
    (assoc s
           :ops ops
           :rest (:rest (last ops))
           :len (+ 18 nlen))))

(defmethod process-op :length [{:keys [rest] :as s}]
  (let [ops (loop [len 0 rest rest final []]
              (let [processed (process (packet rest))
                    nlen (+ len (:len processed))]
                (if (= nlen (:len s))
                  (conj final processed)
                  (recur nlen (:rest processed) (conj final processed)))))]
    (assoc s
           :ops ops
           :len (+ 22 (:len s))
           :rest (:rest (last ops)))))

(defn process-packet [s]
  (-> s parse-hex packet process))

(defn version-counts [s]
  (apply + (map :version (flatten-ops (process-packet s)))))

(defmethod process :literal [s]
  (literal-process s))

(defmethod process :operator [s]
  (-> s operator-process process-op))

(def part-1
  (version-counts final-data))

(defmethod calc-value 4 [e] (:value e))
(defmethod calc-value 0 [e] (apply + (map calc-value (:ops e))))
(defmethod calc-value 1 [e] (apply * (map calc-value (:ops e))))
(defmethod calc-value 2 [e] (apply min (map calc-value (:ops e))))
(defmethod calc-value 3 [e] (apply max (map calc-value (:ops e))))
(defmethod calc-value 5 [e]
  (let [[a b] (map calc-value (:ops e))]
    (if (> a b) 1 0)))
(defmethod calc-value 6 [e]
  (let [[a b] (map calc-value (:ops e))]
    (if (< a b) 1 0)))
(defmethod calc-value 7 [e]
  (let [[a b] (map calc-value (:ops e))]
    (if (= a b) 1 0)))

(def part-2
  (calc-value (process-packet final-data)))

(comment
  (flatten-ops (process-packet "38006F45291200"))
  (flatten-ops (process-packet "EE00D40C823060"))
  (version-counts "8A004A801A8002F478")
  (version-counts "620080001611562C8802118E34")
  (version-counts "C0015000016115A2E0802F182340")
  (version-counts "A0016C880162017C3686B18A3D4780")

  (calc-value (process-packet "C200B40A82"))
  (calc-value (process-packet "04005AC33890"))
  (calc-value (process-packet "D8005AC2A8F0"))

  (try
    (process-packet final-data)
    (catch Exception e
      (println e)
      ))
  (process-packet "A0016C880162017C3686B18A3D4780"))
