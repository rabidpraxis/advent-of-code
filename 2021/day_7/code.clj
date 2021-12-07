(ns day-7.code)

(defn read-input [input]
  (read-string (str "[" input "]")))

(def test-input
  (read-input "16,1,2,0,4,2,7,1,2,14"))

(def final-data
  (read-input (slurp "day_7/input.txt")))

(defn effecient-fuel-used
  ([data]
   (effecient-fuel-used data identity))
  ([data multfn]
   (->> (range (apply max data))
        (map
          (fn [i]
            (reduce
              (fn [coll v]
                (+ coll (multfn (Math/abs (- i v)))))
              0
              data)))
        (apply min))))

(def part-1
  (effecient-fuel-used final-data))

(def fuel
  (mapv (fn [i] (apply + (range 1 i))) (range 1 2000)))

(def part-2
  (effecient-fuel-used final-data #(nth fuel %)))
