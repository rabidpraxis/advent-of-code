(ns day-7.code)

(defn read-input [input]
  (read-string (str "[" input "]")))

(def test-input
  (read-input "16,1,2,0,4,2,7,1,2,14"))

(def final-data
  (read-input (slurp "day_7/input.txt")))

(defn part-1 [data]
  (apply min (map
    (fn [i]
      (reduce
        (fn [coll v]
          (+ coll (Math/abs (- i v))))
        0
        data))
    (set data))))

(part-1 final-data)

(def fuel
  (mapv (fn [i] (apply + (range 1 i))) (range 1 2000)))

(defn part-2 [data]
  (apply min (map
    (fn [i]
      (reduce
        (fn [coll v]
          (+ coll (nth fuel (Math/abs (- i v)))))
        0
        data))
    (range (apply max data)))))

(part-2 final-data)
