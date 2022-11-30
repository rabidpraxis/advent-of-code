(ns advent.day-9
  "This is a masterclass in lazyness"
  (:require
    [clojure.string :as string]
    [advent.answers :as answers]
    [advent.utils :as utils]))

(defn create-game [players]
  {:players players
   :player 1
   :index 1.0
   :marble 1
   :circle (sorted-map 0.0 0 1.0 1)
   :scores (zipmap (map inc (range players)) (repeat 0))})

(defn get-index [circle index]
  (let [[a b] (-> (concat (subseq circle >= index) (cycle circle)) rest keys)]
    (if (< a b)
      (/ (+ a b) 2.0)
      (-> circle rseq ffirst inc))))

(defn get-behind-index [circle index n]
  (-> (concat (rsubseq circle <= index) (cycle (rseq circle)))
      keys (nth n)))

(defn advance [game]
  (let [{:keys [players player index marble circle scores]} game
        marble (inc marble)
        player (inc (mod player players))
        game (assoc game :player player :marble marble)]
    (if (zero? (mod marble 23))
      (let [drop-index (get-behind-index circle index 7)
            scores (update scores player + marble (circle drop-index))
            circle (dissoc circle drop-index)
            index (ffirst (or (subseq circle > drop-index) circle))]
        (assoc game :index index :circle circle :scores scores))
      (let [index (get-index circle index)
            circle (assoc circle index marble)]
        (assoc game :index index :circle circle)))))

(defn nth' [coll n]
  (transduce (drop n) (completing #(reduced %2)) nil coll))

(defn solve [players last-marble]
  (-> (iterate advance (create-game players))
      (nth' last-marble)
      :scores vals
      sort last))
