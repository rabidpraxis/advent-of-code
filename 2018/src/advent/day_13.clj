(ns advent.day-13
  (:require
    [clojure.string :as string]
    [clojure.set :as set]
    [clojure.data :refer [diff]]
    [advent.answers :as answers]
    [advent.utils :as utils]))

;; (def input (utils/get-input-lines "day_13_test"))
;; (def input (utils/get-input-lines "day_13_crash"))
(def input (utils/get-input-lines "day_13"))

(def dirs
  {[0 -1] :up
   [0  1] :down
   [-1 0] :left
   [ 1 0] :right})

(def idirs
  (apply hash-map (mapcat reverse dirs)))

(def starting-vel
  {\v (:down idirs)
   \^ (:up idirs)
   \> (:right idirs)
   \< (:left idirs)})

(def cart-angles
  (apply hash-map (mapcat reverse starting-vel)))

(def field
  {:h (count input)
   :w (count (apply max-key count input))})

(def corner-vel-changes
  {[\/ (:right idirs)] (:up idirs)
   [\/ (:left idirs)]  (:down idirs)
   [\/ (:up idirs)]    (:right idirs)
   [\/ (:down idirs)]  (:left idirs)

   [\\ (:right idirs)] (:down idirs)
   [\\ (:left idirs)]  (:up idirs)
   [\\ (:up idirs)]    (:left idirs)
   [\\ (:down idirs)]  (:right idirs)})


(def coord-mapped-pieces
  (remove (comp #{nil \space} second)
          (for [x (range (:w field))
                y (range (:h field))]
            [[x y] (nth (nth input y) x nil)])))

(defn left [[x y]]
  [y (* -1 x)])

(defn right [[x y]]
  [(* -1 y) x])

(def center identity)

(defn get-intersection-fn [idx]
  (nth [left center right] (mod idx 3)))

(def pieces
  (->> coord-mapped-pieces
       (filter (comp #{\\ \/ \+} second))
       (mapcat identity)
       (apply hash-map)))

(defn add-coords [[x1 y1] [x2 y2]]
  [(+ x1 x2) (+ y1 y2)])

(def carts
  (->> coord-mapped-pieces
       (filter (comp starting-vel second))
       (map-indexed (fn [i [coord piece]]
                      {:coord coord
                       :intersections 0
                       :id i
                       :vel (starting-vel piece)}))))

(defn update-cart [{:keys [coord vel intersections id]}]
  (let [ncoord (add-coords coord vel)
        npiece (pieces ncoord)
        [nintersection nvel] (if (= npiece \+)
                               [(inc intersections) ((get-intersection-fn intersections) vel)]
                               [intersections       (or (corner-vel-changes [npiece vel]) vel)])]
    {:coord ncoord
     :prev-coord coord
     :prev-vel vel
     :id id
     :intersections nintersection
     :vel nvel}))

(defn collision? [cart-1 cart-2]
  (or (= (:coord cart-1) (:coord cart-2))
      (= (:coord cart-1) (:prev-coord cart-2))))

(defn find-collisions [carts]
  (loop [carts (sort-by :prev-coord carts)
         ret []]
    (if-let [cart (first carts)]
      (if-let [found (first (filter (partial collision? cart) (rest carts)))]
        (recur (rest carts) (concat ret [cart found]))
        (recur (rest carts) ret))
      (seq ret))))

(defn tick [{:keys [carts tick collisions]}]
  (let [carts (map update-cart (sort-by :coord carts))
        collisions (find-collisions carts)]
    {:carts carts
     :tick (inc tick)
     :collisions collisions}))

(defn clear-carts-from-collision [data]
  (let [{:keys [collisions carts]} data
        id-set (set (map :id collisions))]
    (assoc data
      :collisions nil
      :carts (remove (comp id-set :id) carts))))

(defn print-first-cart [carts]
  (string/join "," (:coord (first carts))))

(defn part-1 []
  (->> {:carts carts :tick 0 :collisions false}
       (iterate tick)
       (filter :collisions)
       first
       :collisions
       print-first-cart))

(defn part-2 []
  (->> {:carts carts :tick 0 :collisions nil}
       (iterate (comp clear-carts-from-collision tick))
       (filter #(= 1 (count (:carts %))))
       first
       :carts
       print-first-cart))

;;
;; Debugging Below
;;
(defn replace-char [input char coord]
  (let [[x y] coord]
    (update input y #(str (subs % 0 x) char (subs % (inc x) (count %))))))

(defn clean-carts [input carts]
  (reduce
    (fn [input {:keys [coord vel]}]
      (let [replacement (or (and (#{:up :down} (dirs vel)) \|) \-)]
        (replace-char input replacement coord)))
    input
    carts))

(def clean-input (clean-carts input carts))

(defn print-input [input]
  (println (string/join \newline input)))

(defn print-input-with-carts [carts input]
  (print-input
    (reduce
      (fn [input {:keys [coord vel]}]
        (let [replacement (cart-angles vel)]
          (replace-char input replacement coord)))
      input
      carts)))

(comment
  (print-input clean-input)

  (doseq [e (take 25 (iterate tick {:carts carts :tick 0 :collisions nil}))]
    (println e))

  (->
    (iterate (comp clear-carts-from-collision tick) {:carts carts :tick 0 :collisions false})
    (nth 144)
    :carts
    ;; println
    ;; :collisions
    (->> (map :coord))
    ;; (print-input-with-carts clean-input)
    )

  (first (filter (comp seq compare-nth) (range 3230 3500)))
  (seq (compare-nth 11693))
  (seq (compare-nth 144))
  (print-input input)

  (->>
    (iterate tick {:carts carts :tick 0 :collisions false})
    (filter :collisions)
    first
    ;; :tick
    :collisions
    ;; ffirst
    ))
