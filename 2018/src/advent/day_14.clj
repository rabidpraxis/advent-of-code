(ns advent.day-14
  (:require
    [clojure.string :as string]
    [advent.answers :as answers]
    [advent.utils :as utils]))

(def input 290431)

(defn split-number [num]
  (-> num str (string/split #"") (->> (map read-string))))

(def input-vec (vec (split-number 290431)))

(defn run-recipe-vec [[scores elves]]
  (let [to-add (split-number (reduce + (mapv (partial get scores) elves)))
        nscores (apply conj scores to-add)
        nelves (mapv #(mod (+ 1 % (get nscores %)) (count nscores)) elves)]
    [nscores nelves]))

(defn part-1 []
  (->> (iterate run-recipe-vec [[3 7] [0 1]])
    (filter (comp (partial < (+ input 10)) count first))
    ffirst
    ((fn [e] (subvec e input (+ input 10))))
    (string/join "")))

(defn part-2 []
  (->>
    (iterate run-recipe-vec [[3 7] [0 1]])
    (filter (fn [[scores _]]
              (let [sc (count scores)]
                (and (> sc 11) (or (= (subvec scores (- sc 6)) input-vec)
                                   (= (subvec scores (- sc 7) (- sc 1)) input-vec))))))
    first
    ((fn [[scores _]]
       (let [sc (count scores)]
         (- sc
            (if (= (subvec scores (- sc 6)) input-vec)
              6 7)))))))
