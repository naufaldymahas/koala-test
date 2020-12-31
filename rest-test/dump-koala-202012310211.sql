PGDMP                         x            koala    12.2    12.2 "    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    25991    koala    DATABASE     c   CREATE DATABASE koala WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'C' LC_CTYPE = 'C';
    DROP DATABASE koala;
                naufaldymahas    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                naufaldymahas    false            �           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   naufaldymahas    false    3            �            1259    25992 	   customers    TABLE     �  CREATE TABLE public.customers (
    customer_id character varying(64) NOT NULL,
    customer_name character varying(80) NOT NULL,
    email character varying(50) NOT NULL,
    phone_number character varying(20) NOT NULL,
    dob date NOT NULL,
    sex boolean NOT NULL,
    salt character varying(80) NOT NULL,
    password character varying(400) NOT NULL,
    created_date timestamp without time zone NOT NULL
);
    DROP TABLE public.customers;
       public         heap    dev    false    3            �            1259    26062    order_details    TABLE       CREATE TABLE public.order_details (
    order_detail_id character varying(64) NOT NULL,
    order_id character varying(64) NOT NULL,
    product_id character varying(64) NOT NULL,
    qty integer NOT NULL,
    created_date timestamp without time zone NOT NULL
);
 !   DROP TABLE public.order_details;
       public         heap    dev    false    3            �            1259    26079    order_month_seq    SEQUENCE     x   CREATE SEQUENCE public.order_month_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.order_month_seq;
       public          dev    false    3            �            1259    26077    order_ref_seq    SEQUENCE     x   CREATE SEQUENCE public.order_ref_seq
    START WITH 123
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.order_ref_seq;
       public          dev    false    3            �            1259    26081    order_year_seq    SEQUENCE     z   CREATE SEQUENCE public.order_year_seq
    START WITH 2019
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.order_year_seq;
       public          dev    false    3            �            1259    26045    orders    TABLE       CREATE TABLE public.orders (
    order_id character varying(64) NOT NULL,
    customer_id character varying(64) NOT NULL,
    order_number character varying(40) NOT NULL,
    order_date timestamp without time zone NOT NULL,
    payment_method_id character varying(64) NOT NULL
);
    DROP TABLE public.orders;
       public         heap    dev    false    3            �            1259    26004    payment_methods    TABLE     �   CREATE TABLE public.payment_methods (
    payment_method_id character varying(64) NOT NULL,
    method_name character varying(70) NOT NULL,
    code character varying(10) NOT NULL,
    created_date timestamp without time zone NOT NULL
);
 #   DROP TABLE public.payment_methods;
       public         heap    dev    false    3            �            1259    26009    products    TABLE     �   CREATE TABLE public.products (
    product_id character varying(64) NOT NULL,
    product_name character varying(80) NOT NULL,
    base_price bigint NOT NULL,
    created_date timestamp without time zone NOT NULL
);
    DROP TABLE public.products;
       public         heap    dev    false    3            �          0    25992 	   customers 
   TABLE DATA           |   COPY public.customers (customer_id, customer_name, email, phone_number, dob, sex, salt, password, created_date) FROM stdin;
    public          dev    false    202            �          0    26062    order_details 
   TABLE DATA           a   COPY public.order_details (order_detail_id, order_id, product_id, qty, created_date) FROM stdin;
    public          dev    false    206            �          0    26045    orders 
   TABLE DATA           d   COPY public.orders (order_id, customer_id, order_number, order_date, payment_method_id) FROM stdin;
    public          dev    false    205            �          0    26004    payment_methods 
   TABLE DATA           ]   COPY public.payment_methods (payment_method_id, method_name, code, created_date) FROM stdin;
    public          dev    false    203            �          0    26009    products 
   TABLE DATA           V   COPY public.products (product_id, product_name, base_price, created_date) FROM stdin;
    public          dev    false    204            �           0    0    order_month_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.order_month_seq', 12, true);
          public          dev    false    208            �           0    0    order_ref_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.order_ref_seq', 130, true);
          public          dev    false    207            �           0    0    order_year_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.order_year_seq', 2019, false);
          public          dev    false    209                       2606    26066    order_details OrderDetails_pkey 
   CONSTRAINT     l   ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT "OrderDetails_pkey" PRIMARY KEY (order_detail_id);
 K   ALTER TABLE ONLY public.order_details DROP CONSTRAINT "OrderDetails_pkey";
       public            dev    false    206                       2606    26051    orders Orders_order_number_key 
   CONSTRAINT     c   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT "Orders_order_number_key" UNIQUE (order_number);
 J   ALTER TABLE ONLY public.orders DROP CONSTRAINT "Orders_order_number_key";
       public            dev    false    205                       2606    26049    orders Orders_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT "Orders_pkey" PRIMARY KEY (order_id);
 >   ALTER TABLE ONLY public.orders DROP CONSTRAINT "Orders_pkey";
       public            dev    false    205                       2606    26001    customers customers_email_key 
   CONSTRAINT     Y   ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_email_key UNIQUE (email);
 G   ALTER TABLE ONLY public.customers DROP CONSTRAINT customers_email_key;
       public            dev    false    202            
           2606    26003 $   customers customers_phone_number_key 
   CONSTRAINT     g   ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_phone_number_key UNIQUE (phone_number);
 N   ALTER TABLE ONLY public.customers DROP CONSTRAINT customers_phone_number_key;
       public            dev    false    202                       2606    25999    customers customers_pkey 
   CONSTRAINT     _   ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (customer_id);
 B   ALTER TABLE ONLY public.customers DROP CONSTRAINT customers_pkey;
       public            dev    false    202                       2606    26008 #   payment_methods paymentmethods_pkey 
   CONSTRAINT     p   ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT paymentmethods_pkey PRIMARY KEY (payment_method_id);
 M   ALTER TABLE ONLY public.payment_methods DROP CONSTRAINT paymentmethods_pkey;
       public            dev    false    203                       2606    26013    products products_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (product_id);
 @   ALTER TABLE ONLY public.products DROP CONSTRAINT products_pkey;
       public            dev    false    204                       2606    26067 (   order_details OrderDetails_order_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT "OrderDetails_order_id_fkey" FOREIGN KEY (order_id) REFERENCES public.orders(order_id);
 T   ALTER TABLE ONLY public.order_details DROP CONSTRAINT "OrderDetails_order_id_fkey";
       public          dev    false    206    205    3092                       2606    26072 *   order_details OrderDetails_product_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT "OrderDetails_product_id_fkey" FOREIGN KEY (product_id) REFERENCES public.products(product_id);
 V   ALTER TABLE ONLY public.order_details DROP CONSTRAINT "OrderDetails_product_id_fkey";
       public          dev    false    206    3088    204                       2606    26052    orders Orders_customer_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT "Orders_customer_id_fkey" FOREIGN KEY (customer_id) REFERENCES public.customers(customer_id);
 J   ALTER TABLE ONLY public.orders DROP CONSTRAINT "Orders_customer_id_fkey";
       public          dev    false    205    3084    202                       2606    26057 $   orders Orders_payment_method_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT "Orders_payment_method_id_fkey" FOREIGN KEY (payment_method_id) REFERENCES public.payment_methods(payment_method_id);
 P   ALTER TABLE ONLY public.orders DROP CONSTRAINT "Orders_payment_method_id_fkey";
       public          dev    false    3086    205    203            �   �  x�e��N[1��9O�8x.�x"!5��
V�7u���-E���}�:�H+dK��}�G�c"/�cǢ�J��b���'���6ۣ���ҫ}ZԻ�ɧ�3�jt>:�i;����_�����v}�����o׫���R�br��2y�-US�ި�*�$L�Z��>��k!�ة�M��;@G�a	�\D�dV�(��Ī:nB.����zp�Sė�/�O�����#�oֻ�����^]���@LB�0�
D
��d��S, �bN�UCl�C��s5|I��5$�(�N�f�lɕ܃�)�T�e��@y���{
�^Р�.~���W�}�;/�����O��tK�������T����e!�b)���6�$��L�d�ћ�����K�@��k
��w⡍Vbs9��E3U`*�@���R ���$=`x`67��������6��o�g���dպ�Q���� SGIf�i�=�po.��v��<.�,	�p�_�<��|��      �   ?  x��һ�d1P�u� ]��8HH���rw���dPp���L[�g*�����%�q,�ȣ��xL���&��Y�n۞���zBP�)3�e�}������� ����1~����j,�D�wLȣ���Gad<,'Hρ�s-K����}w�I����o�|��,N*�]�jIB踖�c��0oL$��,��y�ö�(��ح"2�M�k�C�^�n��3DS��)���٧�nK�.x�Fw�hL��c=�
�ǁ�fP�ޢ�r�^6tF<�ˢ�!~K�p��Uт*�MxWF���_%����\~��,?������>      �   �   x���;n1 �z�� �|��J�i�ƾ��_)i$B��Е��&��Zͦ��B�y9���FІ���َ�w ����ygd<~®����a|s3w;�R��S:h�Y�a�h2��,�IׂʵeClu����NQ���K�<�o,���&�^�к
�g@w#��I�2��S��E�9.P�{�ZWE�yaɯL�Lۃ�&.���U�^���ն̼��ip�_��<���m�      �   �   x�}�;�A и�^����j3�5\6���8.x{����1�ւt�$������c�M���t�6���ˠE��iki����8�����2�e�XC�.��1Tg.�Ko��>�~V."���YT�7QP��
�Y*���N�+�6����>��ne��a� ��:      �   �   x�}���0�s2������,\�4�� �����¬K�=a���dҌXZ��Y���~zH 0��7�ش.�P�F��p�B#s"u��jF�d���)L����w��uR��I�L6D�[ϣ�U[����>�8_f�%��)�5b     